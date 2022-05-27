interface Tuple {
  subject: {
    type: string;
    name: string;
  };
  relation: string;
  resource: {
    type: string;
    name: string;
  };
}

export default Tuple;
