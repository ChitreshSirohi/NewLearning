import React from "react";
import Color from "./color.component";


export default function ColorList({colors = []}) {
    if (!colors.length) return <div>No color found</div>;
    return (
        <div>
            {colors.map(color => <Color {...color} />)}
        </div>
    )
}