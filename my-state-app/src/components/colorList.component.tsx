import React from "react";
import Color from "./color.component";
import {colorDat} from "./color-app.component";


export default function ColorList({colors=colorDat,onRemoveColor=(id:string)=>{console.log(id)}}) {

    if (!colors.length) return <div>No color found</div>;
    return (
        <div>
            {colors.map(color => <Color onRemove={onRemoveColor} key={color.color} {...color} />)}
        </div>
    )
}