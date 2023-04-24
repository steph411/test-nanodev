import { useSearchParams } from "react-router-dom";

const ConfirmEmailPage = ({}) => {
  const params = useSearchParams();
  console.log({ urlParams: params });
  return <div>Email confirmed</div>;
};

export default ConfirmEmailPage;
