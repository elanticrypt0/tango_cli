package templates

func (t *Templates) ViewsFullSelector() string {

	t.setReplacements()

	template := `
templ Selector$SC$(list []models.$SC$){

<label class="text-gray-700" for="item_id">
    $PL$
    <select name="$SL$_id" id="$SL$_id" class="block px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm w-52 focus:outline-none focus:ring-primary-500 focus:border-primary-500">
    for _, item :=range(list){
        <option value={item.GetIDAsString()}>{item.Name}</option>
    }
    </select>
</label>
}
	`

	return t.Replacements.Replace(template)

}
