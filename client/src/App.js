import React from 'react';
import { render } from 'react-dom';


const App = () => {
  return React.createElement("div", { id: "something-important" }, [
    "Hello World!"
  ]);
};

render(React.createElement(App), document.getElementById("root"));
