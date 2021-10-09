import React from 'react';
import { FaStar } from "react-icons/fa";

const Star = ({ selected = false, onSelect = (f: any) => f }) => (
    <FaStar color={selected ? "red" : "grey"} onClick={onSelect} />
);
const createArray = (length:number) => [...Array(length)];

export default function StarRating({ totalStars = 5, selectedStars=0,onRate}:{totalStars:number,selectedStars:number,onRate:(rating:number)=> void}) {
    return (
        <>
            {createArray(totalStars).map((n, i) => (
                <Star
                    key={i}
                    selected={selectedStars > i}
                    onSelect={() => onRate(i+1)}
                />
            ))}
            <p>
                {selectedStars} of {totalStars} stars
            </p>
        </>
    );
}