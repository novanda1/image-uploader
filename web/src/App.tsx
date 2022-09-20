import { useState } from "react";

function App() {
  return (
    <div className="font-sans min-h-screen w-[100vw] bg-light flex flex-col items-center justify-between px-4">
      <div className="container grow flex justify-center items-center">
        <div className="w-[402px] max-w-full bg-white shadow-[0_4px_12px_0_rgba(0,0,0,0.1)] py-9 px-8 rounded-[12px]">
          <h2 className="font-medium text-lg text-gray2 text-center tracking-[0.03rem]">
            Upload Your Image
          </h2>
          <p className="mt-4 text-[10px] text-gray3 text-center">
            File should be Jpeg, Png,...
          </p>

          <p className="text-center mt-[18px] text-gray4 font-medium text-xs">
            Or
          </p>

          <button className="bg-primary-blue rounded-[8px] py-2 px-4 text-center block mx-auto mt-[29px]">
            Choose a file
          </button>
        </div>
      </div>

      <span className="font-montserrat text-[#A9A9A9] text-center text-sm mb-6">
        created by <strong>novanda1</strong> - devChallanges.io
      </span>
    </div>
  );
}

export default App;
