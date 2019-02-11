/*
 * swaggerpetstore_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package store_pkg


import(
	"strconv"
	"encoding/json"
	"swaggerpetstore_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"swaggerpetstore_lib/apihelper_pkg"
	"swaggerpetstore_lib/configuration_pkg"
)
/*
 * Client structure as interface implementation
 */
type STORE_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Returns a map of status codes to quantities
 * @return	Returns the map[string]int64 response from the API call
 */
func (me *STORE_IMPL) GetInventory () (map[string]int64, error) {
    //the endpoint path uri
    _pathUrl := "/store/inventory"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.SERVER1,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "Authorization" : fmt.Sprintf("Bearer %s", me.config.OAuthAccessToken()),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal map[string]int64
    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Place an order for a pet
 * @param    *models_pkg.Order        body     parameter: Required
 * @return	Returns the *models_pkg.Order response from the API call
 */
func (me *STORE_IMPL) CreatePlaceOrder (
            body *models_pkg.Order) (*models_pkg.Order, error) {
    //the endpoint path uri
    _pathUrl := "/store/order"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.SERVER1,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("Bearer %s", me.config.OAuthAccessToken()),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid Order", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.Order = &models_pkg.Order{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * For valid response try integer IDs with value >= 1 and <= 10. Other values will generated exceptions
 * @param    int64        orderId     parameter: Required
 * @return	Returns the *models_pkg.Order response from the API call
 */
func (me *STORE_IMPL) GetOrderById (
            orderId int64) (*models_pkg.Order, error) {
    //the endpoint path uri
    _pathUrl := "/store/order/{orderId}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "orderId" : orderId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.SERVER1,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("Bearer %s", me.config.OAuthAccessToken()),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid ID supplied", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("Order not found", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.Order = &models_pkg.Order{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * For valid response try integer IDs with positive integer value. Negative or non-integer values will generate API errors
 * @param    int64        orderId     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *STORE_IMPL) DeleteOrder (
            orderId int64) (error) {
    //the endpoint path uri
    _pathUrl := "/store/order/{orderId}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "orderId" : orderId,
    })
    if err != nil {
        //error in template param handling
        return err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.SERVER1,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "Authorization" : fmt.Sprintf("Bearer %s", me.config.OAuthAccessToken()),
    }

    //prepare API request
    _request := unirest.Delete(_queryBuilder, headers, nil)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid ID supplied", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("Order not found", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return err
    }

    //returning the response
    return nil

}

