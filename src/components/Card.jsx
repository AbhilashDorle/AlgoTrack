import React from "react";
import './Card.css'

export default function Card(){

    return(
        <div className="scrollable-card-wrapper">
            <label className="title-label">
                <input placeholder="Revision Title" type="text" className="title-input-tag"></input>
            </label>

            <label className="description-label">
                <textarea placeholder="Description" type="text" className="description-input-tag"></textarea>
            </label>
        </div>
    );
}

