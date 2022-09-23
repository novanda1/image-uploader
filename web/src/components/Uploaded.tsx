import { useEffect } from "react";
import {
  useNavigate,
  useSearchParams,
} from "react-router-dom";
import {
  toast,
  ToastContainer,
} from "react-toastify";
import "react-toastify/dist/ReactToastify.min.css";

const Uploaded: React.FC = () => {
  const [searhParams] = useSearchParams();
  const navigate = useNavigate();

  const name = searhParams.get("name") || "";

  const url = `${import.meta.env.VITE_API_URL}/v1/image/${name}`;
  const copy = () => {
    navigator.clipboard
      .writeText(url)
      .then(() => {
        toast.success(
          "Successfully copied link to your clipboard!",
          {
            position: "bottom-right",
          }
        );
      })
      .catch(() => {
        toast.error("Failed to copy link", {
          position: "bottom-right",
        });
      });
  };

  useEffect(() => {
    if (!name) {
      navigate("/");
    }
  }, []);

  return (
    <div>
      <svg
        className="scale-150 mb-[11px] mx-auto"
        xmlns="http://www.w3.org/2000/svg"
        height="24"
        viewBox="0 0 24 24"
        width="24"
      >
        <path d="M0 0h24v24H0z" fill="none" />
        <path
          fill="#219653"
          d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"
        />
      </svg>

      <h2 className="text-lg text-center mb-[25px]">
        Uploaded Successfully!
      </h2>

      <img
        className="w-full h-auto object-cover rounded-[12px] mb-[25px]"
        alt={name}
        src={url}
      />

      <div className="relative flex gap-[14px] items-center bg-[#F6F8FB] rounded-[13px] border border-[#E0E0E0] py-0.5 pl-[15px] pr-0.5">
        {/* text */}
        <p className="text-xs text-gray2 truncate grow shrink basis-0">
          {url}
        </p>
        {/* btn */}
        <button
          onClick={copy}
          className="px-[18px] py-[9px] w-max bg-[#2F80ED] rounded-[10px] text-xs text-white"
        >
          Copy link
        </button>
      </div>

      <ToastContainer />
    </div>
  );
};

export default Uploaded;
