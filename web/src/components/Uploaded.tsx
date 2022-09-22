import { useParams } from "react-router-dom";

const Uploaded: React.FC = () => {
  const { name } = useParams();

  console.log({ name });

  return <div>uploaded...</div>;
};

export default Uploaded;
