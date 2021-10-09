import React, {useState} from "react";
import colorData from "../data/color-data.json"
import ColorList from "./colorList.component";

export interface IColor {
    id: string,
    title: string,
    color: string,
    rating: number
}

export const colorDat: IColor[] = []
export default function ColorApp() {
    const [colors, setColors] = useState(colorData);
    return <ColorList colors={colors}
                      onRemoveColor={(id: string) => {
                          const newColors = colors.filter(color => color.id !== id);
                          setColors(newColors);
                      }}
                      onRateColor={(id, rating) => {
                          const newColors = colors.map(color => color.id === id ? {...color, rating} : color);
                          setColors(newColors);
                      }}

    />;
}
