package requests

import (
	"fmt"
	"net/url"
	"strings"
)

type Filter struct {
	Field string
	Type  FilterType
	Value any
}

type Query struct {
	Filters   []Filter
	TableName string
}

func NewQuery() *Query {
	query := &Query{}
	return query
}

func (q *Query) AddNewFilter(filter string, filterType FilterType, value any) {
	querFilter := Filter{
		Field: filter,
		Type:  filterType,
		Value: value,
	}
	q.Filters = append(q.Filters, querFilter)
}

func (q *Query) SetTable(tableName string) {
	q.TableName = tableName
}

func (Q *Query) parseUrl(targetUrl string) (url.Values, error) {
	urlParts, err := url.Parse(targetUrl)
	if err != nil {
		return nil, err
	}

	return url.ParseQuery(urlParts.RawQuery)
}

// ParseURLQuery: Will parse url.Values into new query filters
func (q *Query) ParseUrl(targetUrl string) error {
	values, err := q.parseUrl(targetUrl)
	if err != nil {
		return err
	}
	for key, values := range values {
		if len(values) == 0 {
			continue
		}

		// this will return []string. eg: [age gt]
		parts := strings.Split(key, ".")

		if len(parts) != 2 {
			continue
		}

		field := parts[0]
		operator := parts[1]

		var filterType FilterType
		switch operator {
		case "eq":
			filterType = Eq
		case "gt":
			filterType = Gt
		case "gte":
			filterType = Gte
		case "lt":
			filterType = Lt
		case "lte":
			filterType = Lte
		case "neq":
			filterType = Neq
		case "like":
			filterType = Like
		case "ilike":
			filterType = ILike
		// Todo: Add more cases
		default:
			continue
		}

		value := values[0]

		q.AddNewFilter(field, filterType, value)
	}
	return nil
}

func (q *Query) filterCondition(filter Filter, pos int) (string, interface{}, error) {
	switch filter.Type {
	case Eq:
		return fmt.Sprintf("%s = $%d", filter.Field, pos), filter.Value, nil
	case Gt:
		return fmt.Sprintf("%s > $%d", filter.Field, pos), filter.Value, nil
	case Gte:
		return fmt.Sprintf("%s >= $%d", filter.Field, pos), filter.Value, nil
	case Lt:
		return fmt.Sprintf("%s < $%d", filter.Field, pos), filter.Value, nil
	case Lte:
		return fmt.Sprintf("%s <= $%d", filter.Field, pos), filter.Value, nil
	case Neq:
		return fmt.Sprintf("%s != $%d", filter.Field, pos), filter.Value, nil
	case Like:
		return fmt.Sprintf("%s LIKE $%d", filter.Field, pos), "%" + filter.Value.(string) + "%", nil
	case ILike:
		return fmt.Sprintf("%s ILIKE $%d", filter.Field, pos), "%" + filter.Value.(string) + "%", nil
	// Todo: Add more filter here
	default:
		return "", nil, fmt.Errorf("unsupported filter type: %v", filter.Type)
	}
}

func (q *Query) SelectQuery() (string, []any, error) {
	var conditions []string
	var args []any
	posCounter := 1
	for _, filter := range q.Filters {
		condition, arg, err := q.filterCondition(filter, posCounter)
		if err != nil {
			return "", nil, err
		}

		conditions = append(conditions, condition)
		args = append(args, arg)
		posCounter++
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM public.%s", q.TableName)
	if len(conditions) > 0 {
		sqlQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	return sqlQuery, args, nil
}
