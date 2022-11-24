package other

var stringFields map[string]*string
var intFields map[string]*int64
var floatFields map[string]*float64
var boolFields map[string]*bool

func init() {
	stringFields = make(map[string]*string)
	//stringFields[nameField] = nil
	//stringFields[typeField] = nil
	//stringFields[resultField] = nil

	intFields = make(map[string]*int64)
	//intFields[amountField] = nil
	//intFields[countField] = nil

	floatFields = make(map[string]*float64)
	//floatFields[timeField] = nil

	boolFields = make(map[string]*bool)
	//boolFields[enabledField] = nil
}

// string fields
const (
	nameField     = "name"
	typeField     = "type"
	categoryField = "category"
	resultField   = "result"
)

// int fields
const (
	amountField = "amount"
	countField  = "result_count"
)

//float fields
const (
	timeField = "energy_required"
)

// bool fields
const (
	enabledField = "enabled"
)

//difficult objects
const (
	normalField      = "normal"
	expensiveField   = "expensive"
	ingredientsField = "ingredients"
)
