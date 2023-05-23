import { TextField } from "@mui/material";
import { Param } from "../keymap";

const KeyPicker = ({
  param,
  update,
}: {
  param: Param;
  update: (param: Param) => void;
}) => {
  return (
    <TextField
      variant="outlined"
      value={param.keyCode}
      onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
        update({
          number: param.number,
          keyCode: param.keyCode,
        });
      }}
    />
  );
};

export default KeyPicker;
