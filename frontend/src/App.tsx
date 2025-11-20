import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [getOneTimestamp, setGetOneTimestamp] = useState<(Date | null)[]>([null, null])
  const [getAllTimestamp, setGetAllTimestamp] = useState<(Date | null)[]>([null, null])
  const [postTimestamp, setPostTimestamp] = useState<(Date | null)[]>([null, null])
  const [updateTimestamp, setUpdateTimestamp] = useState<(Date | null)[]>([null, null])
  const [deleteTimestamp, setDeleteTimestamp] = useState<(Date | null)[]>([null, null])

  const formatTime = (date: Date | null): string => {
    if (!date) return '';
    
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');
    const seconds = date.getSeconds().toString().padStart(2, '0');
    const milliseconds = date.getMilliseconds().toString().padStart(3, '0');
    
    return `${hours}:${minutes}:${seconds}.${milliseconds}`;
  };

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>

      <div className='card'>
        <button onClick={async () => {
          await fetch('http://localhost:8080/api/v1/items/1', {
            method: 'GET',
          }).then(res => res.json()).then(data => {
            console.log(data);
            setGetOneTimestamp([new Date(data.timestamp), new Date()]);
          });
        }}>
          Fetch Items from Backend
        </button>
        <p>Get Items Server Time: {formatTime(getOneTimestamp[0])}</p>
        <p>Get Items Client Time: {formatTime(getOneTimestamp[1])}</p>
      </div>

      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
