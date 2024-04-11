package queryManager

import (
	"fmt"
	"strings"
)

// type CommonFilterQueryTemplate struct {
// 	baseQuery   string
// 	whereClause string
// }

// func (cfqt CommonFilterQueryTemplate) BuildQuery(searchCondition bool) string {
// 	cfqt.baseQuery = fmt.Sprintf(cfqt.baseQuery)
// 	if !searchCondition {
// 		cfqt.whereClause = ""
// 	}
// 	return fmt.Sprintf("%s %s %s %s", cfqt.baseQuery, cfqt.whereClause, cfqt.groupBy, cfqt.orderBy)
// }

// func BuildSQLQuery(baseQuery string, params ParsedQueryParams, conditions []string) (string, []any) {
// 	var args []any
// 	var whereClauses []string

// 	for _, condition := range conditions {
// 		if value, exists := params[condition]; exists {
// 			switch value.(type) {
// 			case int:
// 				if value.(int) != -1 {
// 					whereClauses = append(whereClauses, fmt.Sprintf("%s = $%d", condition, len(args)+1))
// 					args = append(args, value)
// 				}
// 			case string:
// 				if value.(string) != "" {
// 					whereClauses = append(whereClauses, fmt.Sprintf("%s = $%d", condition, len(args)+1))
// 					args = append(args, value)
// 				}
// 			}
// 		}
// 	}

// 	if len(whereClauses) > 0 {
// 		baseQuery += " WHERE " + strings.Join(whereClauses, " AND ")
// 	}

// 	if limit, exists := params["limit"]; exists {
// 		baseQuery += fmt.Sprintf(" LIMIT $%d", len(args)+1)
// 		args = append(args, limit)
// 	}
// 	if offset, exists := params["offset"]; exists {
// 		baseQuery += fmt.Sprintf(" OFFSET $%d", len(args)+1)
// 		args = append(args, offset)
// 	}

// 	return baseQuery, args
// }

type JoinCondition struct {
	Table  string
	JoinOn string
	// Alias      string
	// ForeignKey string
}

func BuildSQLQuery(baseQuery string, params map[string]any, conditions []string) (string, []any) {
	var args []any
	var whereClauses []string
	// var joinClauses []string

	// for _, join := range joins {
	// 	joinClause := fmt.Sprintf("JOIN %s ON %s", join.Table, join.JoinOn)
	// 	joinClauses = append(joinClauses, joinClause)
	// }

	for _, condition := range conditions {
		if value, exists := params[condition]; exists && value != -1 {
			whereClause := fmt.Sprintf("%s = $%d", condition, len(args)+1)
			whereClauses = append(whereClauses, whereClause)
			args = append(args, value)
		}
	}

	// if len(joinClauses) > 0 {
	// 	baseQuery += " " + strings.Join(joinClauses, " ")
	// }

	if len(whereClauses) > 0 {
		baseQuery += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	if limit, exists := params["limit"]; exists {
		baseQuery += fmt.Sprintf(" LIMIT $%d", len(args)+1)
		args = append(args, limit)
	}
	if offset, exists := params["offset"]; exists {
		baseQuery += fmt.Sprintf(" OFFSET $%d", len(args)+1)
		args = append(args, offset)
	}

	return baseQuery, args
}
