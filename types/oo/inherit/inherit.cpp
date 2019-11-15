#include <stdio.h>

class base {
public:
    virtual ~base(){}
    virtual void bar() = 0;
    void foo() {
        bar();
    }
};

class service1 : public base {
public:
    void bar() {
        printf("service1 bar\r\n");
    }
};

class service2 : public base {
public:
    void bar() {
        printf("service2 bar\r\n");
    }
};

int main() {
    base* bs[2] = {0};
    bs[0] = new service1();
    bs[1] = new service2();
    for (auto b : bs) {
        b->foo();
    }
    getchar();
}