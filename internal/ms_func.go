package internal

type msCond struct {
	Suffix, Condition string
}

func ms_conds() []msCond {
	return []msCond{
		msCond{"_EQ", "="},
		msCond{"_NotEQ", "!="},
		msCond{"_LT", "<"},
		msCond{"_LE", "<="},
		msCond{"_GT", ">"},
		msCond{"_GE", ">="},
	}

}

func ms_str_cond() []msCond {
	return []msCond{
		msCond{"_EQ", "="},
	}

}

func ms_in() []msCond {
	return []msCond{
		msCond{"_In", " IN "},
		msCond{"_NotIn", " NOT IN "},
	}

}

func ms_gen_types() []string {
	return []string{"Deleter", "Updater", "Selector"}
}

func ms_to_slice(typ ...string) []string {
	return typ
}
