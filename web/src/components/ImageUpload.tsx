import { useEffect, useState } from "react";
import Dropzone from "./Dropzone";
import Loading from "./Loading";
import Uploaded from "./Uploaded";

export type State = "idle" | "loading" | "uploaded";

const ImageUpload = () => {
  const [uploadData, setUploadData] = useState({});
  const [state, setState] = useState<State>("idle");
  const [file, setFile] = useState<File | null>();

  const upload = async () => {
    if (!file) return;

    setState("loading");
    let formdata = new FormData();
    formdata.append("name", file.name);
    formdata.append("file", file);

    const data = await fetch("http://localhost:4000/v1/image", {
      method: "post",
      body: formdata,
    })
      .then((res) => res.json())
      .then((r) => {
        setState("uploaded");
        return r;
      })
      .catch((err) => {
        console.log(err);
        setState("idle");
      });

    setUploadData(data);

    setFile(null);
  };

  useEffect(() => {
    if (file) upload();
  }, [file]);

  if (state === "loading") return <Loading />;
  else if (state === "uploaded") return <Uploaded data={uploadData} />;

  return <Dropzone setFile={setFile} />;
};

export default ImageUpload;
