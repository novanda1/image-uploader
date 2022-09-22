import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Dropzone from "./Dropzone";
import Loading from "./Loading";

export type State = "idle" | "loading";

const ImageUpload = () => {
  const [state, setState] =
    useState<State>("idle");
  const [file, setFile] = useState<File | null>();

  const navigate = useNavigate();

  const upload = async () => {
    if (!file) return;

    setState("loading");
    let formdata = new FormData();
    formdata.append("name", file.name);
    formdata.append("file", file);

    await fetch(
      "http://localhost:4000/v1/image",
      {
        method: "post",
        body: formdata,
      }
    )
      .then((res) => res.json())
      .then((r) => {
        navigate({
          pathname: "/image",
          search: `?name=${r.data?.name}`,
        });

        return r;
      })
      .catch((err) => {
        console.log(err);
        setState("idle");
      });

    setFile(null);
  };

  useEffect(() => {
    if (file) upload();
  }, [file]);

  if (state === "loading") return <Loading />;

  return <Dropzone setFile={setFile} />;
};

export default ImageUpload;
