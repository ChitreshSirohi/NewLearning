import StarRating from "./star-rating.component";
import React from "react";
import {FaTrash} from "react-icons/fa"

export default function Color({
                                  id,
                                  title,
                                  color,
                                  rating,
                                  onRemove,
                                  onRate
                              }: { id: string, title: string, color: string, rating: number, onRemove: (f: string) => void, onRate: (f: string, rating: number) => void }) {
    return (
        <section>
            <h1>{title}</h1>
            <button onClick={() => onRemove(id)}>
                <FaTrash/>
            </button>
            <div style={{height: 50, backgroundColor: color}}/>
            <StarRating totalStars={5} selectedStars={rating}
                        onRate={rating => onRate(id,rating)}/>
        </section>
    );

}