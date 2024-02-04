package qforge

type Table struct {
	Name    string
	Columns []*Column
}

func (t *Table) GetSharedColumns(other *Table) []*Column {
	columnsMap := make(map[string]struct{})

	// Populate map with IDs from this table
	for _, col1 := range t.Columns {
		columnsMap[col1.Id] = struct{}{}
	}

	var sharedColumns []*Column

	// Check if columns2 has the same IDs as in columns1
	for _, col2 := range other.Columns {
		if _, exists := columnsMap[col2.Id]; exists {
			sharedColumns = append(sharedColumns, col2)
		}
	}

	return sharedColumns
}
