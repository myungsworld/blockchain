class Human {
    public name : string;
    public age : number;
    public gender : string;
    constructor(name:string,age:number,gender:string){
        this.name = name;
        this.age = age;
        this.gender = gender;
    }
}

const song = new Human("song",27,"male")

const sayHi = (person:Human) => {
    return `Hello ${person.name}, you are ${person.age}, you are a ${person.gender}!!`;
};


console.log(sayHi(song));

export {};