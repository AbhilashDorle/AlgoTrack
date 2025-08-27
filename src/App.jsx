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
      { title: "", description: "" } // start empty
    ]);
  };

  const breakpointColumns = {
    default: 3, // desktop
    1100: 3,    // tablet landscape
    700: 2,     // tablet portrait
    500: 1      // mobile
  };

  

  return (
    <div className='main-div'>
      {!loginFlag && <Login changeFlag = {setLoginFlag}/>}
      {loginFlag && 
        
        <div className="logged-in-container">
          
          <button onClick={addCard} className='addCard-btn'>Add card</button>
              <div className="main-app-div">
                {selectedCard !== null && (
                  <div className="overlay" onClick={() => setSelectedCard(null)}>
                    <div className="expanded-card" onClick={(e) => e.stopPropagation()}>
                      <Card {...cardComponent[selectedCard]}/>
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
                      title={card.title}
                      description={card.description}
                      updateCard={(index, field, value) => {
                        setCardComponent(prev => {
                          const newCards = [...prev];
                          newCards[index][field] = value; // update the field
                          return newCards;
                        });
                      }}
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