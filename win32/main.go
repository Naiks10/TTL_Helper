package main

import ( 
	"fmt" 
	"log"
	"golang.org/x/sys/windows/registry"
	"strconv"
)

func main() {
	// Создание ключа
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\Tcpip\Parameters`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	//Обращение к параметру DefaultTTL
	s, _, err := k.GetIntegerValue("DefaultTTL")
	// Если ошибка есть
	if err != nil {
		log.Fatal(err)
	} 
	// Конвератция из UINT(DWord) в string
	var a string = strconv.FormatUint(s, 10)
	fmt.Printf(" DefaultTTL : %q\n Желаете изменить TTL?\n 1) Да\n 2) Нет\n", a)
	var variant int
	fmt.Scan(&variant)
	switch variant {
	case 1:
		fmt.Println("Значение TTL:")
		var ttl uint32
		fmt.Scan(&ttl)
		k.SetDWordValue("DefaultTTL", ttl)
		k, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\Tcpip6\Parameters`, registry.ALL_ACCESS)
		k.SetDWordValue("DefaultTTL", ttl)
	case 2:
		fmt.Println("Операция отменена")
	default:
		fmt.Println("Ошибка!!!")
	}
	k.Close()
	fmt.Println("Нажмите любую клавишу...")
	var b int
	fmt.Scan(&b)
}