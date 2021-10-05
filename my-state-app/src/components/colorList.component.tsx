import React from "react";
import Color from "./color.component";
import {colorDat} from "./color-app.component";


export default function ColorList({colors=colorDat}) {
    if (!colors.length) return <div>No color found</div>;
    return (
        <div>
            {colors.map(color => <Color {...color} />)}
        </div>
    )
}