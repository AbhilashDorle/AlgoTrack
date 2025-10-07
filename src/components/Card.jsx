import React from "react";
import './Card.css'

export default function Card({ id, title, description, index, updateCard, onClick, isSubmitted, deleteCard }) {

    const addToDb = () => {
        console.log("Submit/Update Button Clicked")
        const endpoint = isSubmitted ? 'http://localhost:8080/update/card' : 'http://localhost:8080/add/card';
        const reqType = isSubmitted ? 'PUT' : 'POST';

        const bodyData = isSubmitted ? { id, title, description } : { title, description };
        console.log(bodyData)
        fetch(endpoint, {
            method: reqType,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(bodyData)
        })
            .then(response => response.json())
            .then(data => {
                console.log("Received from backend:", data);
                // Update card in state with returned id and mark as submitted
                updateCard(index, "id", data.id);
                updateCard(index, "isSubmitted", true);
            })
            .catch(error => console.log(error))
    }

    const handleDelete = async () => {
        try {
            const response = await fetch(`http://localhost:8080/cards/${id}`, {
                method: "DELETE",
            });

            if (!response.ok) {
                throw new Error("Failed to delete card");
            }

            // Call parent deleteCard to update state
            deleteCard(index);
        } catch (err) {
            console.error("Error deleting card:", err);
        }
    };


    return (
        <div className="scrollable-card-wrapper" onClick={onClick}>
            <div className="card-container">
                <div className="btn-container">
                    <button className="submit-btn" onClick={addToDb}>
                        {isSubmitted ? "Update" : "Submit"}
                    </button>
                    <button className="delete-btn"
                        onClick={(e) => {
                            handleDelete();
                            e.stopPropagation();
                        }}>
                        Remove
                    </button>
                </div>
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

