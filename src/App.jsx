import React, { useState } from 'react';
import './App.css';
import Card from './components/Card';
import Login from './components/Login';

function App() {

  const [cardComponent, setCardComponent] = useState([]);
  const [loginFlag, setLoginFlag] = useState(false);

  const addCard = () => {
    setCardComponent(prev => [...prev, {}]);
  }

  return (
    <div className='main-div'>
      {!loginFlag && <Login changeFlag = {setLoginFlag}/>}
      {loginFlag && 
        <div className="logged-in-container">
          <button onClick={addCard} className='addCard-btn'>Add card</button>
            <div className='scrollbar-div'>
              <div className="main-app-div">
                {cardComponent.map(()=> <Card />)}
              </div>
            </div>
        </div>
      }
    </div>
  );
}

export default App;