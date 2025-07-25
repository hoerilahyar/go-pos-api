package errors

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

func ParseMySQLError(err error) error {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1045: // ER_ACCESS_DENIED_ERROR
			return ErrDBConnection
		case 1049: // ER_BAD_DB_ERROR
			return ErrDBConnection
		case 1062: // ER_DUP_ENTRY
			return ErrDuplicateEntry
		case 1146: // ER_NO_SUCH_TABLE
			return ErrRecordNotFound
		case 1054: // ER_BAD_FIELD_ERROR
			return ErrInvalidField
		case 1064: // ER_PARSE_ERROR
			return ErrQuerySyntax
		case 1451: // ER_ROW_IS_REFERENCED
			return ErrForeignKeyViolation
		case 1452: // ER_NO_REFERENCED_ROW
			return ErrForeignKeyViolation
		case 1366: // ER_TRUNCATED_WRONG_VALUE
			return ErrInvalidCharacter
		case 1040: // ER_CON_COUNT_ERROR
			return ErrDBConnection
		case 1216: // ER_NO_REFERENCED_ROW_2
			return ErrForeignKeyViolation
		case 1217: // ER_ROW_IS_REFERENCED_2
			return ErrForeignKeyViolation
		case 2002: // CR_CONNECTION_ERROR
			return ErrDBConnection
		case 2013: // CR_SERVER_LOST
			return ErrDBConnection
		case 1292: // ER_TRUNCATED_WRONG_VALUE_FOR_FIELD
			return ErrInvalidDataType
		case 1364: // ER_NO_DEFAULT_FOR_FIELD
			return ErrMissingRequiredField
		case 1048: // ER_BAD_NULL_ERROR
			return ErrNullValue
		default:
			return ErrInternal
		}
	}
	return err
}
