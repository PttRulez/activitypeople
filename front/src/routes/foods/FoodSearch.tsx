import Autocomplete from "src/components/Autocomplete";
import useAxiosPrivate from "src/hooks/useAxiosPrivate";
import useDebounce from "src/hooks/useDebounce";
import { FoodResponse } from "src/types/food";
import { useQuery } from "@tanstack/react-query";
import { useEffect, useState } from "react";

type Props = {
  reactKey: string;
  onChoose: (food: FoodResponse) => void;
  inputProps: {
    [key: string]: any;
  };
};
const FoodSearch = (props: Props) => {
  const { inputProps, onChoose, reactKey } = props;
  const [foodQuery, setFoodQuery] = useState<string>("");
  const debouncedValue = useDebounce<string>(foodQuery, 500);

  const axios = useAxiosPrivate();

  const { data: foods, refetch } = useQuery({
    queryKey: ["foodsearch", debouncedValue],
    queryFn: async () => {
      const data = await axios.get<FoodResponse[]>("/food/search", {
        params: { q: debouncedValue },
      });
      return data.data;
    },
    enabled: false,
  });

  useEffect(() => {
    if (debouncedValue) {
      refetch();
    }
  }, [debouncedValue, refetch]);

  return (
    <Autocomplete<FoodResponse>
      getOptionLabel={(f) => {
        return f.name;
      }}
      items={foods ?? []}
      inputProps={{
        placeholder: "Введите название продукта",
        key: reactKey,
        ...inputProps,
      }}
      onChoose={onChoose}
      onChange={(e: any) => {
        if (inputProps.onChange) inputProps.onChange(e);
        setFoodQuery(e.target.value);
      }}
    />
  );
};

export default FoodSearch;
