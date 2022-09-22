import { useCallback, useEffect, useRef } from "react";
import image from "../assets/image.svg";

type DropzonProps = {
  setFile: React.Dispatch<React.SetStateAction<File | null | undefined>>;
};

const Dropzone: React.FC<DropzonProps> = ({ setFile }) => {
  const dropzoneRef = useRef<HTMLDivElement>(null);

  const handleDragOver = useCallback(() => {
    dropzoneRef.current?.classList.add("!border-blue-700");
  }, []);

  const handleDragLeave = useCallback(() => {
    dropzoneRef.current?.classList.remove("!border-blue-700");
  }, []);

  const handleDrop = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    e.stopPropagation();

    dropzoneRef.current?.classList.remove("!border-blue-700");
    const { files } = e.dataTransfer;
    setFile(files[0]);
  }, []);

  const onInputChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      if (e.target.files?.length) setFile(e.target.files[0]);
    },
    []
  );

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

  return (
    <>
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

      <p className="text-center mt-[18px] text-gray4 font-medium text-xs">Or</p>

      <label className="relative w-max bg-primary-blue hover:bg-blue-400 rounded-[8px] py-2 px-4 text-sm text-center text-white block mx-auto mt-[29px] cursor-pointer">
        Choose a file
        <input
          className="absolute inset-0 opacity-0 -z-10"
          type="file"
          onChange={onInputChange}
        />
      </label>
    </>
  );
};

export default Dropzone;
