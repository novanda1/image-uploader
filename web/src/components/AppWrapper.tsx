import React, { PropsWithChildren } from "react";

const AppWrapper: React.FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="font-sans min-h-screen w-[100vw] bg-light flex flex-col items-center justify-between px-4">
      <div className="container grow flex justify-center items-center">
        <div className="w-[402px] max-w-full bg-white shadow-[0_4px_12px_0_rgba(0,0,0,0.1)] py-9 px-8 rounded-[12px]">
          {children}
        </div>
      </div>

      <span className="font-montserrat text-[#A9A9A9] text-center text-sm mb-6">
        created by <strong>novanda1</strong> - devChallanges.io
      </span>
    </div>
  );
};

export default AppWrapper;
