import { TextField } from "@mui/material";
import { Parameter } from "../keymap";

const KeyPicker = ({
  param,
  update,
}: {
  param: Parameter;
  update: (param: Parameter) => void;
}) => {
  return (
    <TextField
      variant="outlined"
      value={param.keyCodes}
      onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
        update({
          number: param.number,
          keyCodes: param.keyCodes,
        });
      }}
    />
  );
};

export default KeyPicker;
