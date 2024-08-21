import { useState } from 'react';
import './App.css';
import { usePubSubUiEchoMutation } from './client';

function App() {
  const [topic, setTopic] = useState('')
  const [gcpProjectId, setGCPProjectId] = useState('')
  const [trigger, r] = usePubSubUiEchoMutation()

  function handleEcho() {
    trigger({topicSubscription:{
      gcpProjectId: gcpProjectId,
      pubsubTopicName: topic
    }})
      setTopic('')
      setGCPProjectId('')
  }
  
  return (
    <div className="App">
      <input value={gcpProjectId} onChange={(e) => setGCPProjectId(e.target.value)} />
      <input value={topic} onChange={(e) => setTopic(e.target.value)} />
      <button onClick={() => handleEcho()}>Echo</button>
      <br/>
      {JSON.stringify(r)}
    </div>
  );
}

export default App;
