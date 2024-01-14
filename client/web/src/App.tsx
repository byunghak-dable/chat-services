import "./App.css";
import GoogleOAuth from "./component/google-signin";

function App() {
  return (
    <div className="App">
      <body className="App-header">
        <GoogleOAuth />
      </body>
    </div>
  );
}

export default App;
