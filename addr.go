package innernet

type addr struct{ addr string }

func (a *addr) Network() string { return NetworkName }

func (a *addr) String() string { return a.addr }
