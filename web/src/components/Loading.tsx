import styles from "./Loading.module.css";

const Loading: React.FC = () => {
  return (
    <div>
      <h2 className="text-lg text-gray-2 mb-[30px]">Loading...</h2>
      <div className="relative rounded-[8px] h-[6px] bg-[#F2F2F2]">
        <div
          className={`${styles.Bar} w-[100px] bg-[#2F80ED] rounded-[8px] absolute inset-y-0`}
        ></div>
      </div>
    </div>
  );
};

export default Loading;
