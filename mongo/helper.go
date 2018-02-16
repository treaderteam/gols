package mongo

type SimpleModel struct {
	ModelName string
}

func (m SimpleModel) GetCollectionName() string {
	return m.ModelName
}

// Modificate implements Modificatior for SimpleModel model
func (m *SimpleModel) Modificate() {

}

// GetModel implements Modeller for SimpleModel model
func (m *SimpleModel) GetModel() interface{} {
	return m
}
