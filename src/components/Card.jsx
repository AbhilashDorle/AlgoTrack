import React from "react";
import './Card.css'

export default function Card({ title, description, index, updateCard, onClick }){

    return(
        <div className="scrollable-card-wrapper" onClick={onClick}>
            <div className="card-container">
                <button className="submit-btn">submit</button>
                <div className="input-container">
                    <label className="title-label">
                        <input 
                        placeholder="Revision Title" 
                        type="text" 
                        value={title}
                        className="title-input-tag"
                        onChange={e => updateCard(index, "title", e.target.value)}
                        onClick={(e) => e.stopPropagation()}
                        ></input>
                    </label>

                    <label className="description-label">
                        <textarea 
                            placeholder="Description" 
                            type="text" 
                            value={description}
                            className="description-input-tag"
                            onChange={e => updateCard(index, "description", e.target.value)}
                            onClick={(e) => e.stopPropagation()}
                        ></textarea>
                    </label>
                </div>
            </div>
        </div>
    );
}

