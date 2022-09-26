// Pretty print
flight := Flight{
        Origin: "GLA",
        Destination: "JFK",
        Price: 2000,
}

b, err := json.MarshalIndent(flight, "", "  ")
if err != nil {
    fmt.Println(err)
}
fmt.Print(string(b))
