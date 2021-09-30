import StarRating from "./star-rating.component";
import React from "react";

export default function Color({ title ,color,rating}:{title:string,color:string,rating:number}) {
    return (
    <section>
        <h1>{title}</h1>
        <div style={ {height:50 ,backgroundColor:color}}/>
        <StarRating totalStars={rating} />
    </section>)

}