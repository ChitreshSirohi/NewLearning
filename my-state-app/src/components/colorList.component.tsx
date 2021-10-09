import React from "react";
import Color from "./color.component";
import {colorDat, IColor} from "./color-app.component";


export default function ColorList({colors = colorDat, onRemoveColor, onRateColor}:
                                      { colors: IColor[], onRemoveColor: (id: string) => void, onRateColor: (id: string, rating: number) => void }) {

    if (!colors.length) return <div>No color found</div>;
    return (
        <div>
            {colors.map(color => <Color onRemove={onRemoveColor} onRate={(id,rating)=> onRateColor(id,rating)}
                                        key={color.color} {...color} />)}
        </div>
    )
}