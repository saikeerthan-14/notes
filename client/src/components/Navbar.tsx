import { Box, Button, Container, Flex, Text, useColorMode, useColorModeValue } from "@chakra-ui/react";
import { IoMoon } from "react-icons/io5";
import { LuSun } from "react-icons/lu";

export default function Navbar() {
    const {colorMode, toggleColorMode} = useColorMode()
    return (
        <Container maxW={"900px"}>
            <Box bg={useColorModeValue("gray.100", "gray.900")} px={4} my={4} borderRadius={"5"}>
                <Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
                    {/* LEFT SIDE */}
                    <Flex alignItems={"center"} justifyContent={"center"} gap={3} display={{base: "none", sm: "flex"}}>
                        <img src="/react.png" alt="logo" width={50} height={50}/>
                        <Text fontSize={40}>+</Text>
                        <img src="/go.png" alt="logo" width={50} height={50}/>
                        <Text fontSize={40}>=</Text>
                        <img src="/explode.png" alt="logo" width={50} height={50}/>
                    </Flex>
                    {/* RIGHT SIDE */}
                    <Flex alignItems={"center"} gap={3}>
                        <Text fontSize={"lg"} fontWeight={500}>Daily Tasks</Text>
                        <Button onClick={toggleColorMode}>
                            {colorMode === "light" ? <IoMoon /> : <LuSun size={20} />}
                        </Button>
                    </Flex>
                    
                </Flex>
            </Box>
        </Container>
    )
    
}