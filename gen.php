<?php
function package($name,$content) {
    return "package $name\n$content";
}
$types = [
    'int',
    'uint',
    'byte',
    'rune',
    'string',
    'float32',
    'float64',
    'bool'
];
$map = "
func Map{TYPE1_TITLE}sTo{TYPE2_TITLE}s(cb func (e {TYPE1}) {TYPE2},arr []{TYPE1}) []{TYPE2} {
    ret := make([]{TYPE2},len(arr))
    for i := range arr {
        ret[i] = cb(arr[i])
    }
    return ret
}
";
$every = "
func Every{TYPE1_TITLE}(cb func (e {TYPE1}) bool, {TYPE1}s []{TYPE1}) bool {
    for i := range {TYPE1}s {
        if !cb({TYPE1}s[i]) {
            return false
        }
    }
    return true
}
";
$each = "
func Each{TYPE1_TITLE}(cb func (value {TYPE1}, ind int), {TYPE1}s []{TYPE1}) {
    for i := range {TYPE1}s {
        cb({TYPE1}s[i],i)
    }
}
";
$map_functions = [];
$every_functions = [];
$each_functions = [];
foreach ( $types as $type1 ) {
    foreach ( $types as $type2) {
        if ( $type1 == 'bool' && !in_array($type2,['bool','string','int'])) {
            continue;
        }
        if ( $type2 == 'bool' && !in_array($type1,['bool','string','int'])) {
            continue;
        }
        $actual_map = $map;
        $actual_map = str_replace("{TYPE1}",$type1,$actual_map);
        $actual_map = str_replace("{TYPE2}",$type2,$actual_map);
        $actual_map = str_replace("{TYPE1_TITLE}",ucfirst($type1),$actual_map);
        $actual_map = str_replace("{TYPE2_TITLE}",ucfirst($type2),$actual_map);
        array_push($map_functions,$actual_map);
    }
    $actual_every = $every;
    $actual_every = str_replace("{TYPE1_TITLE}",ucfirst($type1),$actual_every);
    $actual_every = str_replace("{TYPE1}",$type1,$actual_every);
    array_push($every_functions,$actual_every);

    $actual_each = $each;
    $actual_each = str_replace("{TYPE1_TITLE}",ucfirst($type1),$actual_each);
    $actual_each = str_replace("{TYPE1}",$type1,$actual_each);
    array_push($each_functions,$actual_each);
}
file_put_contents("./fpmap.go",package("fplib",join("\n",$map_functions)));
file_put_contents("./fpevery.go",package("fplib",join("\n",$every_functions)));
file_put_contents("./fpeach.go",package("fplib",join("\n",$each_functions)));