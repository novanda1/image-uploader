const Uploaded: React.FC<{ data: any }> = ({ data }) => {
  return (
    <div>
      uploaded...
      {JSON.stringify(data)}
    </div>
  );
};

export default Uploaded;
