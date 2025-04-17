import './App.css'
import Something from './components/something'

 function App() {

    async function getJoke(){
      const res = await fetch("https://api.chucknorris.io/jokes/random") ; 
      const data = await res.json()
      console.log(data.value);
    }
      getJoke();
  return (
    <>
      <h1>This is a updated page</h1>
      <Something/>
    </>
  );
}

export default App
