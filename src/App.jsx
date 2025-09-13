import React, { useState } from 'react';
import './App.css';
import Card from './components/Card';
import Login from './components/Login';
import Masonry from 'react-masonry-css';

function App() {

  const [cardComponent, setCardComponent] = useState([]);
  const [selectedCard, setSelectedCard] = useState(null);
  const [loginFlag, setLoginFlag] = useState(false);

  const addCard = () => {
    setCardComponent(prev => [
      ...prev,
      { id: Date.now(), title: "", description: "", isSubmitted: false } // start empty
    ]);
  };

  const deleteCard = (index) => {
    setCardComponent(prev => prev.filter((_, j) => j !== index));
    setSelectedCard(null); // always close overlay
  };

  const breakpointColumns = {
    default: 3, // desktop
    1100: 3,    // tablet landscape
    700: 2,     // tablet portrait
    500: 1      // mobile
  };



  return (
    <div className='main-div'>
      {!loginFlag && <Login changeFlag={setLoginFlag} />}
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
                  index={i}
                  id={card.id}
                  title={card.title}
                  description={card.description}
                  isSubmitted={card.isSubmitted}
                  updateCard={(index, field, value) => {
                    setCardComponent(prev => {
                      const newCards = [...prev];
                      // create a new object to prevent React input issues
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