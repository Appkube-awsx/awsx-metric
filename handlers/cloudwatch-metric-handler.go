package handlers

import (
	"awsx-metric/log"
	"encoding/json"
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-metric-cli/controller"
	"io"
	"net/http"
)

func GetMetricDataWithMultipleQueries(w http.ResponseWriter, r *http.Request) {
	log.Info("Starting //awsx-metric/metric api")
	w.Header().Set("Content-Type", "application/json")

	//region := r.URL.Query().Get("zone")
	//cloudElementId := r.URL.Query().Get("cloudElementId")
	//cloudElementApiUrl := r.URL.Query().Get("cloudElementApiUrl")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, "Error reading request body:", http.StatusInternalServerError, err)
		return
	}

	// Parse the JSON data into a map
	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		handleError(w, "Error decoding JSON data:", http.StatusBadRequest, err)
		return
	}
	//vaultUrl := r.URL.Query().Get("vaultUrl")

	if requestData["cloudWatchQueries"] == nil {
		log.Error("no cloudwatch query provided")
		http.Error(w, fmt.Sprintf("no cloudwatch query provided"), http.StatusBadRequest)
		return
	}
	cloudWatchQueriesRaw, ok := requestData["cloudWatchQueries"]
	if !ok {
		handleError(w, "Missing cloudWatchQueries in the request", http.StatusBadRequest, nil)
		return
	}
	var stringQueries string

	// Type switch to handle different types of queries
	switch queries := cloudWatchQueriesRaw.(type) {
	case []interface{}:
		// Iterate over slice of queries
		for _, query := range queries {
			switch q := query.(type) {
			case string:
				stringQueries += q
			case map[string]interface{}:
				// Assuming you have a function to convert the map to a string representation
				queryString, err := convertMultipleQueriesMapToString(q)
				if err != nil {
					handleError(w, "Error converting map to string", http.StatusInternalServerError, err)
					return
				}
				stringQueries += queryString
			default:
				handleError(w, "Invalid query format in the slice", http.StatusBadRequest, nil)
				return
			}
		}
	case string:
		// Use the provided string query
		stringQueries = queries
	default:
		handleError(w, "Invalid cloudWatchQueries format", http.StatusBadRequest, nil)
		return
	}

	var region = ""
	if requestData["zone"] != nil {
		region = requestData["zone"].(string)
	}

	var cloudElementId = ""
	if requestData["cloudElementId"] != nil {
		cloudElementId = requestData["cloudElementId"].(string)
	}

	var cloudElementApiUrl = ""
	if requestData["cloudElementApiUrl"] != nil {
		cloudElementApiUrl = requestData["cloudElementApiUrl"].(string)
	}

	if cloudElementId != "" {
		authFlag, clientAuth, err := authenticate.AuthenticateData(cloudElementId, cloudElementApiUrl, "", "", "", region, "", "", "", "")
		if err != nil || !authFlag {
			log.Error(err.Error())
			http.Error(w, fmt.Sprintf("Exception: "+err.Error()), http.StatusInternalServerError)
			return
		}
		result, respErr := controller.GetMetricData(clientAuth, stringQueries)
		if respErr != nil {
			log.Error(respErr.Error())
			http.Error(w, fmt.Sprintf("Exception: "+respErr.Error()), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)
	} else {
		var accessKey = ""
		if requestData["accessKey"] != nil {
			accessKey = requestData["accessKey"].(string)
		}

		var secretKey = ""
		if requestData["secretKey"] != nil {
			secretKey = requestData["secretKey"].(string)
		}

		var crossAccountRoleArn = ""
		if requestData["crossAccountRoleArn"] != nil {
			crossAccountRoleArn = requestData["crossAccountRoleArn"].(string)
		}

		var externalId = ""
		if requestData["externalId"] != nil {
			externalId = requestData["externalId"].(string)
		}

		authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
		if err != nil || !authFlag {
			log.Error(err.Error())
			http.Error(w, fmt.Sprintf("Exception: "+err.Error()), http.StatusInternalServerError)
			return
		}
		result, respErr := controller.GetMetricData(clientAuth, stringQueries)
		if respErr != nil {
			///log.Error(respErr.Error())
			http.Error(w, fmt.Sprintf("Exception: "+respErr.Error()), http.StatusInternalServerError)
			return
		}
		fmt.Println("cli response :::: ", result)
		json.NewEncoder(w).Encode(result)
	}

	log.Info("/awsx-metric/metric completed")

}

func GetMetricDataWithSingleQuery(w http.ResponseWriter, r *http.Request) {
	log.Info("Starting //awsx-metric/metric api")
	w.Header().Set("Content-Type", "application/json")

	//region := r.URL.Query().Get("zone")
	//cloudElementId := r.URL.Query().Get("cloudElementId")
	//cloudElementApiUrl := r.URL.Query().Get("cloudElementApiUrl")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, "Error reading request body:", http.StatusInternalServerError, err)
		return
	}

	// Parse the JSON data into a map
	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		handleError(w, "Error decoding JSON data:", http.StatusBadRequest, err)
		return
	}
	//vaultUrl := r.URL.Query().Get("vaultUrl")

	if requestData["cloudWatchQueries"] == nil {
		log.Error("no cloudwatch query provided")
		http.Error(w, fmt.Sprintf("no cloudwatch query provided"), http.StatusBadRequest)
		return
	}
	cloudWatchQueriesRaw, ok := requestData["cloudWatchQueries"]
	if !ok {
		handleError(w, "Missing cloudWatchQueries in the request", http.StatusBadRequest, nil)
		return
	}
	var stringQueries string

	// Type switch to handle different types of queries
	switch queries := cloudWatchQueriesRaw.(type) {
	case []interface{}:
		// Iterate over slice of queries
		for _, query := range queries {
			switch q := query.(type) {
			case string:
				stringQueries += q
			case map[string]interface{}:
				// Assuming you have a function to convert the map to a string representation
				queryString, err := convertMapToString(q)
				if err != nil {
					handleError(w, "Error converting map to string", http.StatusInternalServerError, err)
					return
				}
				stringQueries += queryString
			default:
				handleError(w, "Invalid query format in the slice", http.StatusBadRequest, nil)
				return
			}
		}
	case string:
		// Use the provided string query
		stringQueries = queries
	default:
		handleError(w, "Invalid cloudWatchQueries format", http.StatusBadRequest, nil)
		return
	}

	var region = ""
	if requestData["zone"] != nil {
		region = requestData["zone"].(string)
	}

	var cloudElementId = ""
	if requestData["cloudElementId"] != nil {
		cloudElementId = requestData["cloudElementId"].(string)
	}

	var cloudElementApiUrl = ""
	if requestData["cloudElementApiUrl"] != nil {
		cloudElementApiUrl = requestData["cloudElementApiUrl"].(string)
	}

	if cloudElementId != "" {
		authFlag, clientAuth, err := authenticate.AuthenticateData(cloudElementId, cloudElementApiUrl, "", "", "", region, "", "", "", "")
		if err != nil || !authFlag {
			log.Error(err.Error())
			http.Error(w, fmt.Sprintf("Exception: "+err.Error()), http.StatusInternalServerError)
			return
		}
		result, respErr := controller.GetMetricDataWithSingleQuery(clientAuth, stringQueries)
		if respErr != nil {
			log.Error(respErr.Error())
			http.Error(w, fmt.Sprintf("Exception: "+respErr.Error()), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)
	} else {
		var accessKey = ""
		if requestData["accessKey"] != nil {
			accessKey = requestData["accessKey"].(string)
		}

		var secretKey = ""
		if requestData["secretKey"] != nil {
			secretKey = requestData["secretKey"].(string)
		}

		var crossAccountRoleArn = ""
		if requestData["crossAccountRoleArn"] != nil {
			crossAccountRoleArn = requestData["crossAccountRoleArn"].(string)
		}

		var externalId = ""
		if requestData["externalId"] != nil {
			externalId = requestData["externalId"].(string)
		}

		authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
		if err != nil || !authFlag {
			log.Error(err.Error())
			http.Error(w, fmt.Sprintf("Exception: "+err.Error()), http.StatusInternalServerError)
			return
		}
		result, respErr := controller.GetMetricDataWithSingleQuery(clientAuth, stringQueries)
		if respErr != nil {
			///log.Error(respErr.Error())
			http.Error(w, fmt.Sprintf("Exception: "+respErr.Error()), http.StatusInternalServerError)
			return
		}
		fmt.Println("cli response :::: ", result)
		json.NewEncoder(w).Encode(result)
	}

	log.Info("/awsx-metric/metric completed")

}

func handleError(w http.ResponseWriter, logMsg string, statusCode int, err error) {
	log.Error(logMsg, err)
	http.Error(w, fmt.Sprintf("Exception: %s", logMsg), statusCode)
}

func convertMapToString(queryMap map[string]interface{}) (string, error) {

	// Convert the array to a JSON string
	jsonBytes, err := json.Marshal(queryMap)
	if err != nil {
		return "", err
	}

	// Convert the JSON bytes to a string
	queryString := string(jsonBytes)

	return queryString, nil
}

func convertMultipleQueriesMapToString(queryMap map[string]interface{}) (string, error) {
	// Convert the map to a JSON string
	var keyValueArray []map[string]interface{}
	for key, value := range queryMap {
		keyValueArray = append(keyValueArray, map[string]interface{}{key: value})
	}

	// Convert the array to a JSON string
	jsonBytes, err := json.Marshal(keyValueArray)
	if err != nil {
		return "", err
	}

	// Convert the JSON bytes to a string
	queryString := string(jsonBytes)

	return queryString, nil
}
