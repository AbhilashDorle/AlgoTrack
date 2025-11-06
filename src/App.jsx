import React, { useState } from 'react';
import './App.css';
import Card from './components/Card';
import Login from './components/Login';
import Masonry from 'react-masonry-css';

function App() {
  const [cardComponent, setCardComponent] = useState([]);
  const [selectedCard, setSelectedCard] = useState(null);
  const [loginFlag, setLoginFlag] = useState(false);
  const [token, setToken] = useState(null)

  const [userId, setUserId] = useState(null); // track logged-in user

  // Add a new empty card
  const addCard = () => {
    setCardComponent(prev => [
      ...prev,
      { id: Date.now(), title: "", description: "", isSubmitted: false }
    ]);
  };

  // Delete card locally
  const deleteCard = (index) => {
    setCardComponent(prev => prev.filter((_, i) => i !== index));
    setSelectedCard(null);
  };

  // Update card fields
  const updateCard = (index, field, value) => {
    setCardComponent(prev => {
      const newCards = [...prev];
      newCards[index] = { ...newCards[index], [field]: value };
      return newCards;
    });

  };

  //Replace with useEffect with user_id in the dependency array
  //This way the system will GET all cards on login
  //Need to handle logic for signup
  const handleGet = async (user_id) => {
    try {
      const response = await fetch(`http://localhost:8080/getcards/${user_id}`, {
        method: "GET",
        headers: {
          "Authorization": `Bearer ${token}`
        },
      });

      if (!response.ok) {
        console.log(response)
        throw new Error("Failed to get cards");
      }

      const data = await response.json();
      console.log(data)
      setCardComponent(data.map(card => ({
        ...card,
        isSubmitted: true // So UI knows these came from DB
      })));
    } catch (err) {
      console.error("Error retrieving cards:", err);
    }
  };

  const handleSuggestion = async (user_id) => {
    try {
      const response = await fetch(`http://localhost:8080/suggestions/${user_id}`, {
        method: "GET",
        headers: {
          "Authorization": `Bearer ${token}`
        },
      });

      if (!response.ok) {
        console.log(response)
        throw new Error("Failed to get suggestions");
      }

      const data = await response.json();
      console.log(data)
      setCardComponent(data.map(card => ({
        ...card,
        isSubmitted: true // So UI knows these came from DB
      })));
    } catch (err) {
      console.error("Error retrieving cards:", err);
    }
  };

  const breakpointColumns = {
    default: 3,
    1100: 3,
    700: 2,
    500: 1
  };

  return (
    <div className='main-div'>
      {!loginFlag &&
        <Login
          setToken={setToken}
          setUserId={setUserId}
          changeFlag={(flag) => {
            setLoginFlag(flag);
          }}
        />
      }
      {/* ******LOGOUT FUNCTIONALITY REMAINS************ */}

      {loginFlag &&
        <div className="logged-in-container">
          <div className='button-div'>
            <button onClick={() => handleSuggestion(userId)} className='addCard-btn'>Help Me Revise</button>
            <div className='db-btn'>
              <button onClick={addCard} className='addCard-btn'>Add New Card</button>
              <button onClick={() => handleGet(userId)} className='addCard-btn'>My Existing Cards</button>
            </div>
          </div>
          <div className="main-app-div">
            {selectedCard !== null && (
              <div className="overlay" onClick={() => setSelectedCard(null)}>
                <div className="expanded-card" onClick={(e) => e.stopPropagation()}>
                  <Card
                    id={cardComponent[selectedCard].id}
                    title={cardComponent[selectedCard].title}
                    description={cardComponent[selectedCard].description}
                    index={selectedCard}
                    token={token}
                    user_id={userId}
                    isSubmitted={cardComponent[selectedCard].isSubmitted}
                    updateCard={updateCard}
                    deleteCard={deleteCard}
                    onClick={() => { }}
                  />
                </div>
              </div>
            )}

            <Masonry
              breakpointCols={breakpointColumns}
              className="my-masonry-grid"
              columnClassName="my-masonry-grid_column"
            >
              {cardComponent.map((card, i) => (
                <Card
                  key={i}
                  id={card.id}
                  user_id={userId}
                  title={card.title}
                  description={card.description}
                  index={i}
                  token={token}
                  isSubmitted={card.isSubmitted}
                  updateCard={updateCard}
                  deleteCard={deleteCard}
                  onClick={() => setSelectedCard(i)}
                />
              ))}
            </Masonry>
          </div>
        </div>
      }
    </div>
  );
}

export default App;
