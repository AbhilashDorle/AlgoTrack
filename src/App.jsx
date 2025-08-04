import React, { useState } from 'react';
import './App.css';
import Card from './components/Card';

function App() {

  const [cardComponent, setCardComponent] = useState([]);

  const addCard = () => {
    setCardComponent(prev => [...prev, {}]);
  }

  return (
    <div className='main-div'>
      <button onClick={addCard} className='addCard-btn'>Add card</button>
      <div className="main-app-div">
        {cardComponent.map(()=> <Card />)}
      </div>
    </div>
  );
}

export default App;