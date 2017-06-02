package servicenowapi

import "errors"

func (con *apiConnection) tableRequest(method htmlMethods, tableName string, params string) error {
	if !con.IsConnected() {
		return errors.New("tableRequest: Connect() must be run before requesting data")
	}

	switch method {
	case get:
	case post:
	case put:
	case patch:
	case delete:
	default:
		return errors.New("tableRequest: Unsupported method")
	}



	return nil
}

func (con *apiConnection) GetTable(tableName string, queryParams []QueryParameter) error {
	var query string
	for _, item := range queryParams {
		var sep string
		query += sep + item.Key + "=" + item.Value
		sep = ","
	}
	con.tableRequest(get, tableName, query)

	return nil
}
