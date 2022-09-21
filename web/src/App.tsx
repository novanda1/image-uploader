import { useCallback, useEffect, useRef, useState } from "react";
import image from "./assets/image.svg";

function App() {
  const [file, setFile] = useState<File | null>(null);
  const dropzoneRef = useRef<HTMLDivElement>(null);

  const handleDragOver = useCallback(() => {
    dropzoneRef.current?.classList.add("border-blue-700");
  }, []);

  const handleDragLeave = useCallback(() => {
    dropzoneRef.current?.classList.remove("border-blue-700");
  }, []);

  const handleDrop = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    e.stopPropagation();

    dropzoneRef.current?.classList.remove("border-blue-700");
    const { files } = e.dataTransfer;
    setFile(files[0]);
  }, []);

  const onInputChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      if (e.target.files?.length) setFile(e.target.files[0]);
    },
    []
  );

  const getBase64 = (file: File) => {
    return new Promise((resolve) => {
      let fileInfo;
      let baseURL = "";
      let reader = new FileReader();

      reader.readAsDataURL(file);

      reader.onload = () => {
        if (!reader?.result) return;
        baseURL = reader.result as string;
        resolve(baseURL);
      };
    });
  };

  const handleUpload = useCallback(async () => {
    if (!file) return;

    const base64 = await getBase64(file);

    await fetch("http://localhost:4000/v1/image/upload", {
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json, application/xml, text/plain, text/html, *.*",
      },
      method: "post",
      body: JSON.stringify({ file: base64 }),
    });

    setFile(null);
  }, [file]);

  useEffect(() => {
    if (dropzoneRef.current) {
      dropzoneRef.current.addEventListener("dragover", handleDragOver);
      dropzoneRef.current.addEventListener("dragleave", handleDragLeave);
      dropzoneRef.current.addEventListener<any>("drop", handleDrop);
    }

    return () => {
      dropzoneRef.current?.removeEventListener("dragover", handleDragOver);
      dropzoneRef.current?.removeEventListener("dragleave", handleDragLeave);
      dropzoneRef.current?.removeEventListener<any>("drop", handleDrop);
    };
  }, []);

  useEffect(() => {
    if (file) {
      handleUpload();
    }
  }, [file, handleUpload]);

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

          <div
            className="relative border-dashed border-2 border-accent-blue bg-soft-blue rounded-[12px] mt-[29px] py-[35px]"
            ref={dropzoneRef}
          >
            <img role="none" src={image} className="mx-auto mb-[53px]" />
            <p className="text-gray4 text-sm text-center">
              Drag & Drop your image here
            </p>
            <input className="opacity-0 absolute inset-0" />
          </div>

          <p className="text-center mt-[18px] text-gray4 font-medium text-xs">
            Or
          </p>

          <label className="relative w-max bg-primary-blue hover:bg-blue-400 rounded-[8px] py-2 px-4 text-sm text-center text-white block mx-auto mt-[29px]">
            Choose a file
            <input
              className="absolute inset-0 opacity-0 cursor-pointer"
              type="file"
              onChange={onInputChange}
            />
          </label>
        </div>
      </div>

      <span className="font-montserrat text-[#A9A9A9] text-center text-sm mb-6">
        created by <strong>novanda1</strong> - devChallanges.io
      </span>
    </div>
  );
}

export default App;
