import React, { useState } from 'react';
import './App.css';
import Card from './components/Card';
import Login from './components/Login';
import Masonry from 'react-masonry-css';

function App() {
  const [cardComponent, setCardComponent] = useState([]);
  const [selectedCard, setSelectedCard] = useState(null);
  const [loginFlag, setLoginFlag] = useState(false);
  const [userId, setUserId] = useState(null); // track logged-in user

  // Add a new empty card
  const addCard = () => {
    setCardComponent(prev => [
      ...prev,
      { id: Date.now(), title: "", description: "", isSubmitted: false }
    ]);
  };

  // Delete a card locally
  const deleteCard = (index) => {
    setCardComponent(prev => prev.filter((_, j) => j !== index));
    setSelectedCard(null);
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
          changeFlag={(flag, id) => { 
            setLoginFlag(flag); 
            setUserId(id); // capture logged-in user ID
          }} 
        />
      }

      {loginFlag &&
        <div className="logged-in-container">
          <button onClick={addCard} className='addCard-btn'>Add card</button>

          <div className="main-app-div">
            {selectedCard !== null && (
              <div className="overlay" onClick={() => setSelectedCard(null)}>
                <div className="expanded-card" onClick={(e) => e.stopPropagation()}>
                  <Card
                    id={cardComponent[selectedCard].id}
                    title={cardComponent[selectedCard].title}
                    description={cardComponent[selectedCard].description}
                    index={selectedCard}
                    isSubmitted={cardComponent[selectedCard].isSubmitted}
                    userId={userId}
                    updateCard={(index, field, value) => {
                      setCardComponent(prev => {
                        const newCards = [...prev];
                        newCards[index] = { ...newCards[index], [field]: value };
                        return newCards;
                      });
                    }}
                    markSubmitted={() => {
                      setCardComponent(prev => {
                        const newCards = [...prev];
                        newCards[selectedCard] = { ...newCards[selectedCard], isSubmitted: true };
                        return newCards;
                      });
                    }}
                    deleteCard={() => deleteCard(selectedCard)}
                    onClick={() => {}}
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
                  title={card.title}
                  description={card.description}
                  index={i}
                  isSubmitted={card.isSubmitted}
                  userId={userId}
                  updateCard={(index, field, value) => {
                    setCardComponent(prev => {
                      const newCards = [...prev];
                      newCards[index] = { ...newCards[index], [field]: value };
                      return newCards;
                    });
                  }}
                  markSubmitted={() => {
                    setCardComponent(prev => {
                      const newCards = [...prev];
                      newCards[i] = { ...newCards[i], isSubmitted: true };
                      return newCards;
                    });
                  }}
                  deleteCard={() => deleteCard(i)}
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
