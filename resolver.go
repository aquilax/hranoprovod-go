package main

func SumMerge(coef float32, left, right *Elements) Elements{
  var result = Elements{};
  for k, v := range *left {
    _, exists := result[k];
    if (exists){
      result[k] += v * coef;
    } else {
      result[k] = v * coef;
    }
  }
  for k, v := range *right{
    _, exists := result[k];
    if (exists){
      result[k] += v * 1;
    } else {
      result[k] = v * 1;
    }
  }
  return result
}

func (db *NodeList) ResolveNode(name string, level int){
  if (level > 9){
    return
  }
  node, exists := (*db)[name]
  if (!exists){
    return
  }
  var tempa = Elements{}

  for sname, value := range node.elements {
    db.ResolveNode(sname, level+1)
    snode, exists := (*db)[sname]
    if (exists) {
      tempa = SumMerge(value, &snode.elements, &tempa)
    } else {
      tm := make(Elements)
      tm[sname] = value;
      tempa = SumMerge(1, &tempa, &tm)
    }
    node.elements = tempa;
    (*db)[name] = node;
  }
}

func (db *NodeList) Resolve(){
  for name, _ := range *db {
    db.ResolveNode(name, 0);
  }
}
