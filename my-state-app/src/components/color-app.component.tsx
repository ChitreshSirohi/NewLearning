import React, {useState} from "react";
import colorData from "../data/color-data.json"
import ColorList from "./colorList.component";

interface color {
    id:string,
    title:string,
    color:string,
    rating:number
}{colorData}
export default function ColorApp(){

    const [colors,setColorData] = useState(colorData);

    return <ColorList colors={colors} />;
}
