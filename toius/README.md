Preload
err = db.Preload("Role").Find(&accounts).Error
access this with the master object, not the nested object

To JSON
a, err := json.Marshal(p.Source) get json byte array
if err != nil {
	return nil, err
}
n := len(a)        //Find the length of the byte array
s := string(a[:n]) //convert to string
fmt.Println(s)