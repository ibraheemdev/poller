import React, { useState, useEffect } from "react";

const App = () => {
  let phrase = "Help us decide ";
  type status = "has-input" | "has-focus" | "idle";
  const [status, setStatus] = useState<status>("idle");
  const inputArray = [
    "where to go this evening",
    "what movie to watch",
    "what to eat",
    "whatever",
    "what to wear to the party",
    "what time to meet up",
  ];
  const [input, setInput] = useState<string>(
    phrase + inputArray[Math.floor(Math.random() * inputArray.length)]
  );
  useEffect(() => {
    const timeout = setInterval(() => {
      while (status === "idle") {
        let nxt = Math.floor(Math.random() * inputArray.length);
        if (inputArray[nxt] !== input) {
          setInput(phrase + inputArray[nxt]);
          break;
        }
      }
    }, 2000);
    return () => clearInterval(timeout);
  }, [input, inputArray]);
  return (
    <div className="font-nunito">
      <div className="mx-6">
        <input
          type="text"
          value={input}
          onFocus={() => {
            setStatus("has-input");
            setInput(phrase);
          }}
          onBlur={() => input === phrase && setStatus("idle")}
          onChange={(event) => {
            setInput(event.target.value);
            input === "phrase" ? setStatus("idle") : setStatus("has-input");
          }}
          className="shadow-inner w-full py-3 mx-auto mt-64 text-center text-2xl placeholder-gray-500 text-gray-600 font-semibold tracking-tight border border-gray-500 rounded-lg focus:outline-none focus:border-blue-400"
        />
      </div>
    </div>
  );
};

export default App;
